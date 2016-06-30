#!/usr/bin/python
#http://stackoverflow.com/a/1010597/2125837

import argparse

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument(
        'integers', metavar='int', type=int, choices=range(10),
         nargs='+', help='an integer in the range 0..9')
    parser.add_argument(
        '--sum', dest='accumulate', action='store_const', const=sum,
        default=max, help='sum the integers (default: find the max)')

    args = parser.parse_args()
    print(args.accumulate(args.integers))
    
