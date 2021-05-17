# How to contribute

1. Fork the [project](https://github.com/romie-gr/romie) at GitHub.
2. Use `git clone` to fetch a local copy to your computer.
3. Use `git remote add` to set two remote targets, one for the original upstream repo and one for our fork.
4. Create a new branch `git checkout -b`
5. Write the code.
6. Install [act](https://github.com/nektos/act) and run the tests: `act -j test` and the linter `act -j lint`.
7. If they pass, then `git commit` and `git push` to your fork remote against your branch.
8. Now you can create a pull-request from your fork remote to the original upstream project against the `master` branch.

