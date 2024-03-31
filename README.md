# WEEK 1 
-------
## GOLANG
1. **Simplicity and Readability** > Go's syntax is clean and easy to understand
2. **Support for concurrency** >  Go has built-in support for concurrency through go routines and channels.
3. **Static Typing** > Type checking is performed at compile-time
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
    - AST represents the syntatic structure of the input
    - Parser is **responsible to indicate violations** of the langauges's syntax rules and report to the user
---------
