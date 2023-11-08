import ctypes
library = ctypes.CDLL('./lib.so')

library.helloWorld.restype = None
library.helloWorld.argtypes = []

_comp = library.compare
_comp.restype = ctypes.c_char_p
_comp.argtypes = [ctypes.POINTER(ctypes.c_ubyte), ctypes.c_int,
                  ctypes.POINTER(ctypes.c_ubyte), ctypes.c_int]


def hello_world():
    library.helloWorld()


def compare(a: bytes, b: bytes) -> float:
    a_array = (ctypes.c_ubyte * len(a))(*a)
    b_array = (ctypes.c_ubyte * len(b))(*b)
    raw_result = _comp(a_array, len(a), b_array, len(b))
    res = ctypes.cast(raw_result, ctypes.c_char_p).value.decode('utf-8')

    if res.startswith("err:"):
        raise Exception(res)
    num = float(res.split(":")[1])
    return num
