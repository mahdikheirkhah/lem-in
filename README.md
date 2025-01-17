# lem-in

## Objectives

- This project is a digital representation of an ant farm simulation. The goal is to:

- Develop a program called lem-in that reads input from a file describing ants and their colony.

- Calculate and display the quickest path for ants to move through the colony from a start room (##start) to an end room (##end).

- Handle various edge cases, including invalid input formats, unconnected rooms, and circular links.

- Adhere to Go programming language best practices.


## Project Structure

```
    lem-in/
    ├── main.go               
    ├── go.mod
    ├── README.md
    ├── test/
    │   ├── lem-in_test.go
    │   ├── test1.txt
    │   ├── test2.txt
    │   ├── test3.txt
    │   ├── test4.txt  
    │   ├── test5.txt
    │   ├── test6.txt
    │   └── ValidFile.txt
    ├── utils/
    │   ├── checkContent.go
    │   ├── createGraph.go
    │   ├── makeRoom.go
    │   ├── makeTunnel.go
    │   ├── readFromArgs.go    
    │   └── structs.go
    ├── fileHandler/
    │   └── read.go
    ├── errorHandler/
    │    └── checkError.go
    └──  examples/
```

## Usage
### Requirements

Install Go to run and develop the project

### How to Run

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/mahdikheirkhah/lem-in.git
   ```
2. Navigate to the project directory:
   ```bash
   cd lem-in
   ```

3. Save the input describing the colony into a text file, e.g., ant_farm.txt. you can see some examples [here](https://github.com/01-edu/public/tree/master/subjects/lem-in/examples). Or use examples in example folder.
    

4. Run the program with the command::

   ```bash
   go run . examples\example00.txt

   ```
    Or put any file you like as an argument for the programme.
### Examples of Output
#### Example 1

    ```bash
    $ go run . test0.txt
    3
    ##start
    1 23 3
    2 16 7
    3 16 3
    4 16 5
    5 9 3
    6 1 5
    7 4 8
    ##end
    0 9 5
    0-4
    0-6
    1-3
    4-3
    5-2
    3-5
    4-2
    2-1
    7-6
    7-2
    7-4
    6-5

    L1-3 L2-2
    L1-4 L2-5 L3-3
    L1-0 L2-6 L3-4
    L2-0 L3-0
    ```

#### Example 2

    ```bash
    $ go run . test1.txt
    3
    ##start
    0 1 0
    ##end
    1 5 0
    2 9 0
    3 13 0
    0-2
    2-3
    3-1

    L1-2
    L1-3 L2-2
    L1-1 L2-3 L3-2
    L2-1 L3-3
    L3-1
    ```
---

## Instructions

1. Define Rooms and Tunnels:

    -  A room is defined as name coord_x coord_y.
    -  A tunnel is defined as name1-name2.
2. Follow Room Naming Rules:
    -  Names cannot start with L or # and cannot contain spaces.
3. Simulation Rules:

    - Ants start at ##start and aim to reach ##end.

    - Each room can contain only one ant, except ##start and ##end.

    - Each tunnel can be used only once per turn.

    - Ants must avoid traffic jams and use the shortest paths.


4. Error Handling:

    - The program validates input and outputs error messages for invalid data formats or logical issues.

5. Testing:

    - Create test files to validate different edge cases and scenarios.

---

## Authors

- Parisa Rahimi
- Fatemeh Kheirkhah
- Mohammad mahdi Kheirkhah
- Majid Rouhani
---
