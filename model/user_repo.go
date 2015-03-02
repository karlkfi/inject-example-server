package model

type UserRepo interface {
    FindUser(id string) (user User, found bool)
    FindUserByName(name string) (user User, found bool)
    All() UserList
}

type userRepo struct {
    users map[string]User
}

func NewUserRepo(users []User) UserRepo {
    return &userRepo{users: listToMap(users)}
}

func (r *userRepo) FindUser(id string) (User, bool) {
    for id, user := range r.users {
        if user.ID == id {
            return user, true
        }
    }
    return User{}, false
}

func (r *userRepo) FindUserByName(name string) (User, bool) {
    for _, user := range r.users {
        if user.Name == name {
            return user, true
        }
    }
    return User{}, false
}

func (r *userRepo) All() UserList {
    return UserList{Users: mapToList(r.users)}
}

func mapToList(userMap map[string]User) []User {
    userList := make([]User, 0, len(userMap))
    for _, user := range userMap {
        userList = append(userList, user)
    }
    return userList
}

func listToMap(userList []User) map[string]User {
    userMap := make(map[string]User, len(userList))
    for _, user := range userList {
        userMap[user.ID] = user
    }
    return userMap
}