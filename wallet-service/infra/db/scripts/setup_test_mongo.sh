
echo "Setting up MongoDB replica set for testing..."

echo "Cleaning up existing containers..."
docker stop test-mongo &>/dev/null || true
docker rm test-mongo &>/dev/null || true
docker network rm test-mongo-network &>/dev/null || true

echo "Creating test network..."
docker network create test-mongo-network

echo "Starting MongoDB container with replica set configuration..."
docker run -d --name test-mongo \
  --network test-mongo-network \
  -p 27018:27017 \
  mongo:7.0 \
  --replSet rs0 --bind_ip_all

echo "Waiting for MongoDB to start..."
sleep 3

echo "Initializing replica set..."
docker exec test-mongo mongosh --eval '
  rs.initiate({
    _id: "rs0",
    members: [{ _id: 0, host: "localhost:27017" }]
  })
'

echo "Waiting for replica set to stabilize..."
sleep 2

echo "Creating test database and user..."
docker exec test-mongo mongosh --eval '
  rs.status();
  db = db.getSiblingDB("admin");
  db.createUser({
    user: "walletuser",
    pwd: "walletpassword",
    roles: [{ role: "root", db: "admin" }]
  });
  db = db.getSiblingDB("wallet");
  db.createCollection("wallets");
'

echo "MongoDB replica set is ready for testing"
echo "Connection URI: mongodb://walletuser:walletpassword@localhost:27018/wallet?authSource=admin&replicaSet=rs0"