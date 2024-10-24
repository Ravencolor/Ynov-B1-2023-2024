class Watch:
    def __init__(self, hour, minute):
        self.hour = hour
        self.minute = minute

    def clone(self):
        # Retourne une nouvelle instance de Watch avec les mêmes heures et minutes
        return Watch(self.hour, self.minute)

    def advance(self, minutes=1):
        # Avance l'heure de la montre de 'minutes' minutes
        self.minute += minutes
        while self.minute >= 60:
            self.hour += 1
            self.minute -= 60
            if self.hour >= 24:
                self.hour = 0

    def get_time_string(self):
        # Retourne l'heure sous forme de chaîne de caractères
        return f"{self.hour}h{self.minute}"


class Person:
    def __init__(self, name, watch=None):
        self.name = name
        self.watch = watch

    def wear(self, watch):
        # Permet à la personne de porter une montre si elle n'en porte pas déjà une
        if self.watch is None:
            self.watch = watch

    def remove(self):
        # Enlève la montre de la personne
        self.watch = None

    def ask_time(self, person):
        # Demande l'heure à une autre personne
        if person.watch is not None:
            return person.watch.get_time_string()
        else:
            return("la personne n'a pas de montre")


# Création d'une montre avec l'heure 10:10
w = Watch(10, 10)
print(f"Création d'une montre : {w.get_time_string()}")

# Clonage de la montre
w2 = w.clone()
print(f"Clonage de la montre : {w2.get_time_string()}")

# Avancement de la montre de 30 minutes
w.advance(30)
print(f"Avancement de la montre : {w.get_time_string()}")

# Création d'une personne nommée Tintin
tintin = Person("Tintin")
print(f"Création de la personne : {tintin.name}")

# Création d'une personne nommée Dupond avec une montre clonée
dupond = Person("Dupond", w2)
print(f"Création de la personne : {dupond.name}")

# Tintin porte une montre
tintin.wear(w)
print(f"{tintin.name} porte une montre : {tintin.watch.get_time_string()}")

# Tintin enlève sa montre
tintin.remove()
print(f"{tintin.name} a enlevé sa montre")

# Tintin demande l'heure à Dupond
time = tintin.ask_time(dupond)
print(f"{tintin.name} demande l'heure à {dupond.name} , il est {time}")

# Dupond porte une montre
dupond.wear(w2)
print(f"{dupond.name} porte une montre : {dupond.watch.get_time_string()}")

