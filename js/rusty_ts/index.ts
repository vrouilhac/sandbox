type None = null | undefined;
type Some<T> = NonNullable<T>;
type Option<T> = None | Some<T>;

const Some = <T>(value: any): Option<T> => value;
const None: None = null;

const myval = Some(3);

const match = <T>(value: Option<T>) => {
  return (value: any) => {
    return value;
  };
};

let x = match(myval)(
  (() => {
    switch(myval) {
      case None:
	return 1;
      default:
	return myval;
    }
  })()
);

// console.log(myval + 3);
//
console.log(x);
