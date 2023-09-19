import sys

def add_newline_after_braces(input_file):
    try:
        with open(input_file, 'r') as f:
            content = f.read()
    except FileNotFoundError:
        print(f"Error: File '{input_file}' not found.")
        return

    output_content = ''
    indent_level = 0

    for char in content:
        if char == '{':
            output_content += char
            indent_level += 1
        elif char == '}':
            output_content += '}\n' + '  ' * (indent_level - 1)  # Add newline and proper indentation
            indent_level -= 1
        else:
            output_content += char

    print(output_content)

if __name__ == "__main__":
    add_newline_after_braces(sys.argv[1])

