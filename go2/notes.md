2. intro to variables
   if variable name starts with capital letters it is exported and can be seen by other packages
   if not capital letter then it's local to only that package
   declared vars(without initialization) with var have zero value

3. path separator
   path package is util package for working with url paths
   to discard variable return from function we can use blank identifier \_

4. short declaration
   it cannot be used outside of the function. instead use var. also var can be
   used for grouping variable declaration
   use declaration for: init value unknown, package scoped var (outside functions), for grouping variable for better readability
   use short declaration for: if init val is known, concise code, for redeclaration
   what is shadowing?
5. type conversion
   it's simple: type(value) -> boom
   type conversion does not change original value

6. os package
   os.Args stores arguments in slice
   len() function is for slice length
