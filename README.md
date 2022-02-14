![Github Actions](https://github.com/zkmiyavi/invader-sim/actions/workflows/test.yml/badge.svg)
![Github Actions](https://github.com/zkmiyavi/invader-sim/actions/workflows/build.yml/badge.svg)



# Invader Sim

This project implements an alien invasion simulatore as a CLI. This repo is temporary and will be made private after assessment.

## Objective

The objective laid out is to build an alien invasion simulator with constraints laid out in [prompt.txt]()
    
---


## Prompt Assumptions

Reading through the prompt I make the following assumptions:
```
Mad aliens are about to invade the earth and you are tasked with simulating the invasion. 
You are given a map containing the names of cities in the non-existent world of X. The map
is in a file, with one city per line. The city name is first, followed by 1-4 directions 
(north, south, east, or west). Each one represents a road to another city that lies in that 
direction. 

For example: 
Foo north=Bar west=Baz south=Qu-ux 
Bar south=Foo west=Bee 
```
* there are two types of cities: cities declared explicitly in the file(e.g. Foo and Bar) and those declared implicitly (Baz, Qu-ux, Bee and Baz)
* an explicitly declared city can contain up to 4 city connections/links 
* city names do not contain spaces
* input files contain only valid connections:
> Foo north=Bar (new line) Bar south=Foo OR Bar is VALID

> Foo north=Bar (new line) Bar south=Baz is NOT VALID


```
The city and each of the pairs are separated by a single space, and the directions are 
separated from their respective cities with an equals (=) sign. You should create N aliens, 
where N is specified as a command-line argument. These aliens start out at random places on 
the map, and wander around randomly, following links. Each iteration, the aliens can travel 
in any of the directions leading out of a city. In our example above, an alien that starts at 
Foo can go north to Bar, west to Baz, or south to Qu-ux. When two aliens end up in the same 
place, they fight, and in the process kill each other and destroy the city. When a city is 
destroyed, it is removed from the map, and so are any roads that lead into or out of it. 
```
* N aliens can occupy the map. There is no constraint that the number of cities c must be greater than the number of aliens a.
* The initial random placement of aliens into the city counts as the first step. I.e. cities can be destroyed in the first step.
* Making this assumption because if we don't treat this initial phase as the first otherwise you could, for example, initialize 30 aliens onto a map that contains 3 cities and end up with a > 2 aliens in each city, which breaks the constraint that when two aliens end up in the same city, the aliens and city are destroyed.

```
In our example above, if Bar were destroyed the map would now be something like: 
Foo west=Baz south=Qu-ux 
```
* connections between cities must be stored in some state.

```
Once a city is destroyed, aliens can no longer travel to or through it. This may lead to aliens
getting "trapped". You should create a program that reads in the world map, creates N aliens, 
and unleashes them. The program should run until all the aliens have been destroyed, or each 
alien has moved at least 10,000 times. When two aliens fight, print out a message like:

Bar has been destroyed by alien 10 and alien 34!
```
* Aliens must move sequentially, and conflicts must be detected as they occure, otherwise aliens could travel to cities that already contain 2 aliens.
* We need to identify individual aliens, not simply the number of aliens in any city at a given time.

```
(If you want to give them names, you may, but it is not required.) Once the program has finished, 
it should print out whatever is left of the world in the same format as the input file. 
```
* Aliens don't contain much state. Simply creating a slice of integers will give us unique aliens.
* Must maintain some state of all cities and connections.
