db.createUser({
    user: "localUser",
    pwd: "localPassword",
    roles: [{
        role: "readWrite",
        db: "localDb"
    }]
})