class Ingredient:
    def __init__(self, name, state, unit, quantity):
        # Initialisation des attributs de l'ingrédient
        self.name = name
        self.state = state
        self.unit = unit
        self.quantity = quantity

    def equals(self, ingredient):
        # Vérifie si deux ingrédients sont égaux en comparant leur nom et état
        return self.name == ingredient.name and self.state == ingredient.state


class CookableIngredient(Ingredient):
    def __init__(self, name, state, unit, quantity):
        # Initialisation de l'ingrédient cuisinable en utilisant le constructeur de la classe parente
        super().__init__(name, state, unit, quantity)

    def cook(self):
        # Change l'état de l'ingrédient à "cuit"
        self.state = "cooked"


class CutableIngredient(Ingredient):
    def __init__(self, name, state, unit, quantity):
        # Initialisation de l'ingrédient découpable en utilisant le constructeur de la classe parente
        super().__init__(name, state, unit, quantity)

    def cut(self):
        # Change l'état de l'ingrédient à "découpé"
        self.state = "cut"


class Dish:
    def __init__(self, name, lst):
        # Initialisation des attributs du plat
        self.name = name
        self.list = lst

    def equals(self, dish):
        # Vérifie si deux plats sont égaux en comparant les ingrédients un par un
        for i in range(len(self.list)):
            if not self.list[i].equals(dish.list[i]):
                return False
        return True

# Ingrédients
choucroute = CookableIngredient("choucroute", "raw", "g", "500")
lard = CookableIngredient("lard", "whole cooked", "g", "150")
saucisse = CookableIngredient("saucisse", "whole cooked", "piece", "2")
carotte = CutableIngredient("Carotte", "whole", "piece", "2")

# Plats
plat1 = Dish("choucroute", [choucroute, lard, saucisse])

# Cuisiner la choucroute et afficher son état
choucroute.cook()
print(choucroute.state)

# Découper la carotte et afficher son état
carotte.cut()
print(carotte.state)
