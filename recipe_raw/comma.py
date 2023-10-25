import json

# Input file name
input_file = "recipe_raw.txt"

# Output file name
output_file = "recipe_raw.txt"


# Read the input text file
with open(input_file, 'r') as infile:
    data = infile.read()

# Split the input data into separate JSON-like objects
objects = data.split("\n{")

# Add a comma to the end of each object except the last one
for i in range(len(objects) - 1):
    objects[i] = "{" + objects[i] + ","

# Join the modified objects back into a single string
result = "\n".join(objects)

# Write the modified content to the output file
with open(output_file, 'w') as outfile:
    outfile.write(result)

print(f"Commas added to {input_file} and saved to {output_file}")
