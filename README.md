# Folder-Manager

A command-line application written in Go that uses topic modeling to intelligently organize, search, and manage files on your local system. Designed for developers, researchers, and anyone handling large volumes of unstructured files, this tool blends manual control with automated file categorization based on file content.

---

## Features

- **Organize Files Automatically**  
  Uses **topic modeling** and **embedding-based models** (like T5 or SBERT) to classify and move files into predefined folders like `Math`, `NLP`, `ML`, `Work`, etc.

- **Search Files by Content**  
  Perform semantic search through your files using context-aware queries.

- **Manual Move Support**  
  Manually move files using CLI commands when needed.

- **Command History and Undo**  
  View history of file operations and undo recent changes easily.

- **Lightweight & Offline**  
  Fully local and CLI-native—no dependency on external APIs or web apps.

## Demo
**History**
![history output](https://github.com/user-attachments/assets/a018a185-f3dd-437e-bdb8-a92a8d08c9a4)

**Search**
![search output](https://github.com/user-attachments/assets/7574068a-bb38-4b69-852d-935306cba649)

**Organize**
![organize output](https://github.com/user-attachments/assets/a50cc423-4b34-4478-980c-7dfa492d057f)

**Undo**
![undo output](https://github.com/user-attachments/assets/d44e7d22-5852-427e-9134-c331de3137b7)

**Move**
![move output](https://github.com/user-attachments/assets/50eb8f7f-2602-444f-b2f8-2967776351bd)
