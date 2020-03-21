#!/usr/bin/python3
# Product of Array Except Self
given = [4, 5, 1, 8, 2]
products_to_left = [1]
products_to_right = [1]

for i in range(len(given)):
    if i == len(given) - 1:
        break
    products_to_left.append(given[i] * products_to_left[i])

for i in range(len(given)):
    if i == len(given) - 1:
        break
    inserted_value = (
        given[len(given) - (i + 1)]
        * products_to_right[len(products_to_right) - (i + 1)]
    )
    products_to_right = [inserted_value] + products_to_right  # prepend

print(given)
print(products_to_left)
print(products_to_right)
print(
    [products_to_left[i] * products_to_right[i] for i in range(len(products_to_left))]
)
