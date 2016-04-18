#!/usr/bin/python

from optparse import OptionParser
import hashlib

class Uniquify:

    def __init__(self, input, output, verbose):
        self.input = input
        self.output = output
        self.verbose = verbose
        self.hashlist = set() # https://wiki.python.org/moin/TimeComplexity

    @staticmethod
    def hashKey(key):
        """
            If each line is long average, use hashed key (to save memory)
            else, return key directly
        """
        # return key
        return hashlib.md5(key).hexdigest()

    def execute(self):
        try:
            out = open(self.output, 'w')
            with open(self.input) as f: # optimized for large file
                for line in f:
                    # unify line delimiter
                    line = line.replace('\r\n', '\n')
                    line = line.replace('\r', '\n')
                    
                    hashedKey = Uniquify.hashKey(line)
                    if hashedKey in self.hashlist:
                        if self.verbose:
                            print 'Skip on duplicate line'
                    else:
                        self.hashlist.add(hashedKey)
                        out.write(line)
                        if self.verbose:
                            print 'Write to %s: %s' % (self.output, line)
        except:
            pass
        finally:
            out.close()


if __name__ == '__main__':
    usage = """
    %prog [--help|-h]
    %prog --file=<filename> --output=<output-filename> [--verbose]"""
    parser = OptionParser(usage = usage)
    parser.add_option('--file', dest = 'input', help = 'input filename', metavar = 'filename')
    parser.add_option('--output', dest = 'output', help = 'input filename', metavar = 'filename')
    parser.add_option('--verbose', dest = 'verbose', help = 'verbose message', action = 'store_true', default = False)
    (opt, args) = parser.parse_args()
    if not opt.input or not opt.output:
        parser.error('Invalid arguments')

    uni = Uniquify(opt.input, opt.output, opt.verbose)
    uni.execute()