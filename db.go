package main

type User struct {
	Id   int64
	Name string
}

type DB struct {
	users []User
}

func (db *DB) create(username string) User {
	var user User = User{Id: int64(len(db.users)), Name: username}
	db.users = append(db.users, user)
	return user
}

func (db *DB) update(username string, new_username string) User {
	for i, user := range db.users {
		if user.Name == username {
			db.users[i].Name = new_username
		}
	}
	return User{}
}

func (db *DB) delete(username string) {
	for i, user := range db.users {
		if user.Name == username {
			db.users = append(db.users[:i], db.users[i+1:]...)
			return
		}
	}
}

func (db *DB) get(id int64) User {
	for _, user := range db.users {
		if user.Id == id {
			return user
		}
	}
	return User{}
}

func (db *DB) getAll() []User {
	return db.users
}