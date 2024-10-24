import random
from abc import ABC, abstractmethod

# Classe abstraite représentant un personnage
class Character(ABC):
    def __init__(self, name, position, hitbox):
        self.hp = random.randint(50, 100)  # Points de vie aléatoires
        self.mp = random.randint(0, 50)    # Points de magie aléatoires
        self.name = name                   # Nom du personnage
        self.position = position           # Position du personnage
        self.hitbox = hitbox               # Hitbox du personnage

    @abstractmethod
    def special_ability(self):
        pass  # Méthode abstraite pour la capacité spéciale

# Classe représentant un héros, dérivée de Character
class Hero(Character):
    count = 0  # Compteur de héros
    def __init__(self, name, position, hitbox):
        super().__init__(name, position, hitbox)
        self.level = 1  # Niveau initial du héros
        Hero.count += 1  # Incrémentation du compteur de héros

# Classe représentant un Tank, dérivée de Hero
class Tank(Hero):
    def special_ability(self):
        return "Shield"  # Capacité spéciale du Tank

# Classe représentant un Mage, dérivée de Hero
class Mage(Hero):
    def special_ability(self):
        return "Magic"  # Capacité spéciale du Mage

# Classe représentant un Healer, dérivée de Hero
class Healer(Hero):
    def special_ability(self):
        return "Heal"  # Capacité spéciale du Healer

# Classe représentant un Warrior, dérivée de Hero
class Warrior(Hero):
    def __init__(self, name, position, hitbox):
        super().__init__(name, position, hitbox)
        self.mp = 0  # Les Warriors n'ont pas de points de magie

    def special_ability(self):
        return "Attack"  # Capacité spéciale du Warrior

# Classe représentant un monstre, dérivée de Character
class Monster(Character):
    count = 0  # Compteur de monstres
    def __init__(self, name, position, hitbox):
        super().__init__(name, position, hitbox)
        Monster.count += 1  # Incrémentation du compteur de monstres

# Classe représentant un Minion, dérivée de Monster
class Minion(Monster):
    def special_ability(self):
        return "Swarm"  # Capacité spéciale du Minion

# Classe représentant un Buldozer, dérivée de Monster
class Buldozer(Monster):
    def special_ability(self):
        return "Crush"  # Capacité spéciale du Buldozer


