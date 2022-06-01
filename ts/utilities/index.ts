interface Person {
  name: string;
  age: number;
  country: string;
  internetId: User;
}

interface User {
  username: string;
  email: string;
  age: number;
  country: string;
}

const user: User = {
  username: "Vincent",
  email: "vincent@gmail.com",
  age: 23,
  country: "FR",
};

const selectData = (user: User): Pick<User, "username" | "email"> => {
  return {
    username: user.username,
    email: user.email,
  };
};

console.log({ se: selectData(user) });
