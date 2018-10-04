// https://playground.ponylang.io/?gist=7f638892da817882c26958797cbe734d
actor Main
  new create(env: Env) =>
    env.out.print("Hello, world!")
