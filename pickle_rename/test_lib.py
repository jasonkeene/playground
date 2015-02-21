

class Foo(object):
    def __init__(self, name):
        self.name = name

    def __repr__(self):
        return "%s at %s instance of %s.%s" % (self.name, id(self),
                                               type(self).__module__,
                                               type(self).__name__)
