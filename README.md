# gfx
grep for file extension

This is a simple Go program that searches for files with specified extensions in a directory and its subdirectories.

## Usage

1. Ensure you have Go installed on your machine.
2. Clone the repository:

    ```bash
    git clone https://github.com/Bamorph/gfx.git
    cd gfx
    go build
    ```

3. Run the program with the desired file extensions:

    ```bash
    gfx <file_extension1> [<file_extension2> ...]
    ```

   Replace `<file_extension1>`, `<file_extension2>`, etc., with the file extensions you want to search for (e.g., `zip`, `tar`).

   Example:

    ```bash
    go run main.go zip tar
    ```

4. The program will recursively search for files with the specified extensions in the current directory and its subdirectories.

## Notes

- The program uses regular expressions to match file names, so you can use regular expressions as file extensions.
