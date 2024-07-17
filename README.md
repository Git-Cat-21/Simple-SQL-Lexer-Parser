# WEEK 1 
-------
## GOLANG 
1. Go is an open source programming language that google developed.
2. Widely used in the development of web applications, cloud and networking services.
3. **Simplicity and Readability** > Go's syntax is clean and easy to understand
4. **Support for concurrency** >  Go has built-in support for concurrency through go routines and channels.
5. **Static Typing** > Type checking is performed at compile-time
6. **Performance** > Its compilation to machine code makes it fast,  This makes it suitable for building high-performance applications, including web servers and distributed systems.
----------
# WEEK 2
-------
## LEXER AND PARSER
### They collectively form the core components of the typical compiler and interpreter 
1. **LEXER**  
    - The **inital stage** of a parser whose primary function is to break down a sequence of characters into smaller units called ***TOKENS***.
    - They are the **building blocks** that the parser will use to understand the structure and meaning of the input.
2. **PARSER** 
    - Takes the tokens produced by lexer and interprets the structure according to the syntax rules. 
    - It arranges these tokens into a **hierarchial structure** called the ***Abstract Syntax Tree*** (AST)
    - AST represents the syntatic structure of the input.
    - Parser is **responsible to indicate violations** of the langauges's syntax rules and report to the user
---------
# WEEK 3
## SQL LEXER-PARSER 
--------
1. **Main Function**
    - Prompts the user to enter a simple SQL query.
    - Calls the *lexer* function ***to tokenize the input***.
    - Prints the individual tokens.
    - Calls the ***parseSelectStatement*** function to **parse the SQL** Query.
2. **Lexer Function**
    - Iterates through the input string character by character.
    - Collects the characters into a current string until it encounter a whitespace character.
    - When a whitespace character is encountered, it ***appends the current string as a token to the tokens slice***.
    - The *appendToken* function is used ***to determine the type of Token*** and then add it to the tokens slice.
3. **ParseSelectStatement Function**
    - ***Iterates*** through the tokens and ***Initilizes the SelectStatment struct to store the parsed information***.
    - If the *token type* is ***not recognized***, it returns an *error*.

   

