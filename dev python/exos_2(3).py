# Classe de base pour les dispositifs électroniques
class Dispositif:
    def __init__(self, marque, modele):
        self.marque = marque
        self.modele = modele

# Classe pour les ordinateurs portables, hérite de Dispositif
class OrdinateurPortable(Dispositif):
    def __init__(self, marque, modele, taille_ecran):
        super().__init__(marque, modele)
        self.taille_ecran = taille_ecran

# Classe pour les ordinateurs fixes, hérite de Dispositif
class OrdinateurFixe(Dispositif):
    def __init__(self, marque, modele, configuration):
        super().__init__(marque, modele)
        self.configuration = configuration

# Classe pour les consoles de jeu, hérite de Dispositif
class ConsoleJeu(Dispositif):
    def __init__(self, marque, modele, type_console):
        super().__init__(marque, modele)
        self.type_console = type_console

# Classe pour les consoles Switch, hérite de ConsoleJeu
class Switch(ConsoleJeu):
    def __init__(self, marque, modele):
        super().__init__(marque, modele, "Switch")

# Classe pour les consoles Xbox, hérite de ConsoleJeu
class Xbox(ConsoleJeu):
    def __init__(self, marque, modele):
        super().__init__(marque, modele, "Xbox")

# Classe pour les consoles Playstation, hérite de ConsoleJeu
class Playstation(ConsoleJeu):
    def __init__(self, marque, modele):
        super().__init__(marque, modele, "Playstation")

# Classe de base pour les véhicules
class Vehicule:
    def __init__(self, marque, modele):
        self.marque = marque
        self.modele = modele

# Classe pour les voitures, hérite de Vehicule
class Voiture(Vehicule):
    def __init__(self, marque, modele, nb_portes):
        super().__init__(marque, modele)
        self.nb_portes = nb_portes

# Classe pour les camions, hérite de Vehicule
class Camion(Vehicule):
    def __init__(self, marque, modele, capacite_charge):
        super().__init__(marque, modele)
        self.capacite_charge = capacite_charge

# Classe pour les vélos, hérite de Vehicule
class Velo(Vehicule):
    def __init__(self, marque, modele, type_velo):
        super().__init__(marque, modele)
        self.type_velo = type_velo

# Classe pour les trottinettes, hérite de Vehicule
class Trottinette(Vehicule):
    def __init__(self, marque, modele):
        super().__init__(marque, modele)
