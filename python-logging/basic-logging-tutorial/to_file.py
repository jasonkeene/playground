import logging

logging.basicConfig(filename="to_file.log", filemode='w', level=logging.DEBUG)
logging.debug("In depth debug info here!")
logging.info("Just some handy info :)")
logging.warning("Danger Will Robinson")
