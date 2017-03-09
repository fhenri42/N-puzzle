NAME=n-puzzle
SRC = src/astar.go \
			src/finder.go \
			src/heuristique.go \
			src/main.go \
			src/parcing.go \
			src/rotation.go \
			src/solvabiliter.go \

all: $(NAME)

$(NAME): $(SRC)
	go build -o  n-puzzle $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
