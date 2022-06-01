var user = {
    username: "Vincent",
    email: "vincent@gmail.com",
    age: 23,
    country: "FR"
};
var selectData = function (user) {
    return {
        username: user.username,
        email: user.email
    };
};
console.log({ se: selectData(user) });
