class Book:
    def __init__(self, nb_pages, author, isbn, marked_Page):
        self.nb_pages = nb_pages  # Nombre de pages du livre
        self.author = author  # Auteur du livre
        self.isbn = isbn  # Numéro ISBN du livre
        self.marked_Page = marked_Page  # Page marquée dans le livre

    def mark_page(self, page):
        # Marque une page si elle est dans les limites du nombre de pages
        if page < self.nb_pages:
            self.marked_Page = page
        else:
            print("Page number is out of range")  # Affiche un message si la page est hors limite


class Library:
    def __init__(self, books):
        self.books = books  # Liste des livres dans la bibliothèque

    def list_authors(self):
        # Retourne une liste des auteurs des livres dans la bibliothèque
        authors = [book.author for book in self.books]
        return authors
    
    def sum_of_marked_pages(self):
        # Calcule la somme des pages marquées de tous les livres
        return sum([book.marked_Page for book in self.books])


# Création de quelques livres
book1 = Book(100, "Author1", "123456789", 50)
book2 = Book(200, "Author2", "987654321", 100)
book3 = Book(300, "Author3", "123456789", 150)

# Création d'une bibliothèque avec les livres
library = Library([book1, book2, book3])

print("\n")

# Affiche la liste des auteurs
print(library.list_authors())
# Affiche la somme des pages marquées
print(library.sum_of_marked_pages())

# Marque une nouvelle page dans le premier livre
book1.mark_page(101)

print("\n")

# Affiche les pages marquées de chaque livre
print(book1.marked_Page)
print(book2.marked_Page)
print(book3.marked_Page)

print("\n")

# Affiche les pages marquées des livres dans la bibliothèque
print(library.books[0].marked_Page)
print(library.books[1].marked_Page)
print(library.books[2].marked_Page)
