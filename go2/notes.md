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

7. (10.) raw string literals
   if string is between double quotes, it is string literals, if string is between backticks then it's raw string literals. raw string literals can
   contain multiple lines of strings
   string literals is interpreted. raw string literals are not interpreted

8. (11.) len() function doesn't count how many characters in
   string but counts the bytes. string literals are made of
   code points aka runes and runes are made up of bytes so
   each character can be 1 bytes or up to 4 bytes. since
   len() function doesn't count how many characters in string, we need to use
   some other package to count how many characters in string. unicode/utf8 package has RuneCountInString function and it returns the number of characters in string
9. (14.) iota
   iota is a built-in constant generator which generates ever increasing numbers
   when constants are declared in batch we can use iota. it starts with 0 and increases by one
   for succeeding constants. we can also use blank identifier to skip some values
10. (17.) printf
    we need to add \n for printf. it doesn't add new line like println
    %s for string
    %q for string, shows with quotes
    %d for integer
    %f for float -> %.0f means precision, %.1f, %.2f ...
    %t for bool
    %T for integer, float64, bool, string
    %v for any value
    there is some called argument indexing. imagine we have use multiple verbs in printf
    but we only use 2 variables. for printf to work we have to provide equal number of verbs
    but in this case since many verbs point to same variable we can use argument indexes to call them.
    btw indexes start with 1, not with 0. keep that in mind
