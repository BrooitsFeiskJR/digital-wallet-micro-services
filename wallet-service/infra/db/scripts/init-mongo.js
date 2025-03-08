print("Starting MongoDB initialization script...");

try {
  rs.initiate({
    _id: "rs0",
    members: [{ _id: 0, host: "mongo:27017" }]
  });
  print("Replica set initialized successfully");
} catch (e) {
  print("Error initializing replica set (may already be initialized): " + e);
}

sleep(5000);

try {
  const status = rs.status();
  print("Replica set status: " + JSON.stringify(status));
} catch (e) {
  print("Error checking replica set status: " + e);
}

db = db.getSiblingDB('admin');

const users = db.getUsers();
let userExists = false;
for (let i = 0; i < users.users.length; i++) {
  if (users.users[i].user === 'walletuser') {
    userExists = true;
    break;
  }
}

if (!userExists) {
  db.createUser({
    user: 'walletuser',
    pwd: 'walletpassword',
    roles: [
      { role: 'readWrite', db: 'wallet-service' },
      { role: 'dbAdmin', db: 'wallet-service' }
    ]
  });
  
  print('User walletuser created successfully');
} else {
  print('User walletuser already exists');
}

db = db.getSiblingDB('wallet-service');
db.createCollection('wallets');
print('Initialized wallet-service database with collections');