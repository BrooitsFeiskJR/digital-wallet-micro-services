db = db.getSiblingDB('admin');

// Check if user already exists
const users = db.getUsers();
let userExists = false;
for (let i = 0; i < users.users.length; i++) {
  if (users.users[i].user === 'walletuser') {
    userExists = true;
    break;
  }
}

// Create user if it doesn't exist
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

// Initialize wallet-service database
db = db.getSiblingDB('wallet-service');
db.createCollection('wallets');
print('Initialized wallet-service database with collections');