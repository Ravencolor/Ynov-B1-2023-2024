class CoffeeShop:

    def __init__(self, name, stock_quantity):
        # Initialisation de la classe avec le nom du café et la quantité de stock
        self.name = name
        self.stock_quantity = stock_quantity
        self.state = None

    def get_state(self):
        # Détermine l'état du stock (True si stock > 5, sinon False)
        if self.stock_quantity > 5:
            self.state = True
        else:
            self.state = False
        return self.state

    def restock(self):
        # Ajoute 10 unités au stock et affiche un message de réapprovisionnement
        self.stock_quantity += 10
        print("Restocked\n")
        return self.stock_quantity

    def serve(self):
        # Sert le client si le stock est suffisant, sinon réapprovisionne et sert à nouveau
        if self.get_state() == True:
            print("Welcome to " + self.name + "!\n")
            number_of_coffee = int(input("How many coffees would you like?\n"))
            self.take_commande(number_of_coffee)
        else:
            print("Sorry, we are restocking. Please come back later.\n")
            print("Restocking...\n")
            print("Many hours later... \n")
            self.restock()
            self.serve()

    def take_commande(self, number_of_coffee):
        # Prend la commande et met à jour le stock en conséquence
        if self.stock_quantity >= number_of_coffee:
            print("Here are your " + str(number_of_coffee) + " coffees. Enjoy!\n")
            self.stock_quantity -= number_of_coffee
            print("Remaining stock: " + str(self.stock_quantity) + "\n")
        else:
            print("Sorry, we don't have enough coffee. Please come back later.\n")

# Création d'une instance de CoffeeShop et appel de la méthode serve
coffee_shop = CoffeeShop("Coffee and Code", 100)
coffee_shop.serve()
