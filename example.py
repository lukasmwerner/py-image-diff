import imagediff

# with open("a.png", "rb") as a:
#     with open("b.png", "rb") as b:
#         difference = library.compare(a.read(), b.read())
#         print(difference)

with open("data_chart.png", 'rb') as a:
    with open("data_chart_other.png", 'rb') as b:
        difference = imagediff.compare(a.read(), b.read())
        print(difference)
