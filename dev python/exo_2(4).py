class StringUtils:
    # Définition des propriétés statiques pour une nouvelle ligne et une tabulation
    new_line = "\n"
    tab = "\t"

    @staticmethod
    def get_first(str):
        # Retourne le premier caractère de la chaîne si elle n'est pas vide, sinon None
        return str[0] if str else None

    @staticmethod
    def get_last(str):
        # Retourne le dernier caractère de la chaîne si elle n'est pas vide, sinon None
        return str[-1] if str else None

    @staticmethod
    def get_substring(str, first, last):
        # Retourne une sous-chaîne de 'first' à 'last' si la chaîne n'est pas vide et 'first' est inférieur à 'last', sinon None
        return str[first:last] if str and first < last else None
    
# Utilisation des propriétés statiques
print(StringUtils.new_line)  # Affiche une nouvelle ligne
print(StringUtils.tab)       # Affiche une tabulation

# Appel des méthodes statiques
print(StringUtils.get_first("Hello"))  # Affiche le premier caractère de "Hello"
print(StringUtils.get_last("Hello"))   # Affiche le dernier caractère de "Hello"
print(StringUtils.get_substring("Hello", 1, 4))  # Affiche la sous-chaîne de "Hello" de l'index 1 à 4 (exclus)