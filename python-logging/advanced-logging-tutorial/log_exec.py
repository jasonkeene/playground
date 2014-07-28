import logging


logging.basicConfig(filename="log_exec.log", filemode="w")
logger = logging.getLogger(__name__)
print logger is logging.getLogger(__name__)


try:
    foobar
except NameError:
    logger.warn("OMG! A NameError was raised!", exc_info=True)
    logger.error("OMG! A NameError was raised!")
    logger.exception("OMG! A NameError was raised!")
