import logging

import mylib


def main():
    logging.basicConfig(filename="myapp.log", level=logging.INFO,
                        format="%(asctime)s | %(levelname)s | %(message)s",
                        datefmt="%Y-%m-%d %I:%M:%S %p",
                        filemode="w")
    logging.info("Started myapp!")
    mylib.do_something()
    logging.warn("%s before you %s", "Leap", "look!")
    logging.info("Done!")


if __name__ == "__main__":
    main()
