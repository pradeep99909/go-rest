package main

type User struct {
	Id   int
	Name string
}

type DB struct {
	users []User
}

func (db *DB) create(username string) User {
	var user User = User{Id: len(db.users), Name: username}
	db.users = append(db.users, user)
	return user
}

func (db *DB) update(username string, new_username string) User {
	for _, user := range db.users {
		if user.Name == username {
			user.Name = new_username
			return user
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

func (db *DB) get(id int) User {
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