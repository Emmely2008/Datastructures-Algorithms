## 1. Compulsory algorithm Assignment - Experiments and Sorting.
This is the hand-in paper for Datastructure & Algorithms course. It consists of 3 bullets points (Subjects) which needs attention. This paper will include results, comments to results and the essential code.

I've chossen to make all experiments and programs in "Python". My codes can be run an executed online.

## Birthday Experiment
#### Description 
Write a program that takes an integer N from the command line and generate a random sequence of integers between 0 and N – 1. Run experiments to validate the hypothesis that the number of integers generated before the first repeated value is found is ~√(pi)N/2.

#### Execute Program Online

jdoodle.com/a/nln

#### Results & comments:

Results to this problem is as follows:
```
With n=30 runs=1000.0: got result=7.467 hypothesis=6.86468424648
With n=80 runs=1000.0: got result=11.792 hypothesis=11.2099824328
```
This seems very close, the diffence is: ~0.671 and might therefor be true. However i wanted to try see what would happen if you limited towards infinity to see if the expected formular would get closer to actual tests. So new results show from if N was 1000 and not 365.

The difference is ~0.656. Which gets higher than previus. But it stays about the same offset for any value put in to N. This indicates it is very close to the actual formular.

The big O notation is ofcause the formular. O(√(pi)N/2) But worst case scenario, it could also be O(N). It could potentially, go through the whole available valuespace before it finds a duplicate. Which is N. 

#### Essential Code:

```python
def birthdayProblem(n):
    i = 0
    a = []
    while( i < n ):
      chance = math.floor(random.random()*n);
      for j in  a:
        if(j==chance):
          i = n
          break      
      a.append(chance)
      i += 1
    return a
```


## Coupon collector problem.
#### Description
Generating random integers as in the previous exercise run experiments to validate the hypothesis that the number of integers generated before all possible values are generated is ~N HN.
#### Execute Program Online

jdoodle.com/a/nlk

#### Results & comments

```
With n=6 runs=1000.0 result=14.606
With n=20 runs=1000.0 result=71.939
```

#### Essential Code

The fucntion to run this test is as follows:
```python
def couponCollectorProblem(n):
	a = range(n)
	i = 0
	while( len(a) > 0 ):
	  chance = math.floor(random.random()*n);
	  i += 1
	  if chance in a:
		try:
		  a.remove(chance)
		except:
		  pass
	return i;   
``` 

## Deck sort
#### Description
Explain how you would put a deck of cards in order by suit (in the
order spades, hearts, clubs, diamonds) and by rank within each suit, with the restriction
that the cards must be laid out face down in a row, and the only allowed operations are
to check the values of two cards and to exchange two cards (keeping them face down).

#### Execute Program Online

jdoodle.com/a/nll

#### Results & Comments:

```
 deck
[0:Spades1, 1:Spades2, 2:Spades3, 3:Spades4, 4:Spades5, 5:Spades6, 6:Spades7, 7:Spades8, 8:Spades9, 9:Spades10, 10:SpadesJ, 11:SpadesQ, 12:SpadesK, 13:Hearts1, 14:Hearts2, 15:Hearts3, 16:Hearts4, 17:Hearts5, 18:Hearts6, 19:Hearts7, 20:Hearts8, 21:Hearts9, 22:Hearts10, 23:HeartsJ, 24:HeartsQ, 25:HeartsK, 26:Clubs1, 27:Clubs2, 28:Clubs3, 29:Clubs4, 30:Clubs5, 31:Clubs6, 32:Clubs7, 33:Clubs8, 34:Clubs9, 35:Clubs10, 36:ClubsJ, 37:ClubsQ, 38:ClubsK, 39:Diamonds1, 40:Diamonds2, 41:Diamonds3, 42:Diamonds4, 43:Diamonds5, 44:Diamonds6, 45:Diamonds7, 46:Diamonds8, 47:Diamonds9, 48:Diamonds10, 49:DiamondsJ, 50:DiamondsQ, 51:DiamondsK]
Shuffled deck
[5:Spades6, 38:ClubsK, 26:Clubs1, 35:Clubs10, 43:Diamonds5, 27:Clubs2, 36:ClubsJ, 17:Hearts5, 11:SpadesQ, 19:Hearts7, 24:HeartsQ, 44:Diamonds6, 12:SpadesK, 4:Spades5, 6:Spades7, 15:Hearts3, 13:Hearts1, 21:Hearts9, 39:Diamonds1, 20:Hearts8, 25:HeartsK, 22:Hearts10, 23:HeartsJ, 7:Spades8, 16:Hearts4, 3:Spades4, 48:Diamonds10, 28:Clubs3, 51:DiamondsK, 50:DiamondsQ, 14:Hearts2, 18:Hearts6, 42:Diamonds4, 34:Clubs9, 2:Spades3, 29:Clubs4, 47:Diamonds9, 32:Clubs7, 46:Diamonds8, 9:Spades10, 8:Spades9, 37:ClubsQ, 31:Clubs6, 30:Clubs5, 10:SpadesJ, 1:Spades2, 49:DiamondsJ, 40:Diamonds2, 45:Diamonds7, 33:Clubs8, 41:Diamonds3, 0:Spades1]
Sorted deck
[0:Spades1, 1:Spades2, 2:Spades3, 3:Spades4, 4:Spades5, 5:Spades6, 6:Spades7, 7:Spades8, 8:Spades9, 9:Spades10, 10:SpadesJ, 11:SpadesQ, 12:SpadesK, 13:Hearts1, 14:Hearts2, 15:Hearts3, 16:Hearts4, 17:Hearts5, 18:Hearts6, 19:Hearts7, 20:Hearts8, 21:Hearts9, 22:Hearts10, 23:HeartsJ, 24:HeartsQ, 25:HeartsK, 26:Clubs1, 27:Clubs2, 28:Clubs3, 29:Clubs4, 30:Clubs5, 31:Clubs6, 32:Clubs7, 33:Clubs8, 34:Clubs9, 35:Clubs10, 36:ClubsJ, 37:ClubsQ, 38:ClubsK, 39:Diamonds1, 40:Diamonds2, 41:Diamonds3, 42:Diamonds4, 43:Diamonds5, 44:Diamonds6, 45:Diamonds7, 46:Diamonds8, 47:Diamonds9, 48:Diamonds10, 49:DiamondsJ, 50:DiamondsQ, 51:DiamondsK]
``` 

#### Essential Code: 
```python
  
class Card:
  def __init__(self, name, value):
    self.name = name
    self.value = value 
  def __repr__(self):
    return str(self.value) + ":" + self.name
  def getValue(self):
    return self.value
  def compareTo(self, that):
    BEFORE = 1
    EQUAL = 0
    AFTER = -1
    if(self.getValue() == that.getValue()):
      return EQUAL
    if(self.getValue()< that.getValue()):
      return BEFORE
    if(self.getValue() > that.getValue()):
      return AFTER   


def swap(arr, i, pos):
  temp = arr[i]
  arr[i] = arr[pos]
  arr[pos] = temp
  return arr

	  
def orderCards(list):
  
  for i in range(len(list)):
    smallest_pos=i #pick up a card

    for j in range(i,len(list)):
      #do we have a smaller card? the swap    
      if(list[smallest_pos].compareTo(list[j])==-1):
        smallest_pos=j;

        
       
    list=swap(list,i,smallest_pos)
    
  return list
  
```
