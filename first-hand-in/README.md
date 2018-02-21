## 1. Compulsory algorithm Assignment - Experiments and Sorting.

This hand-in is also available on GitHub. https://github.com/Emmely2008/Datastructures-Algorithms/tree/master/first-hand-in

This is the hand-in paper for Datastructure & Algorithms course. It consists of 3 bullets points (Subjects) which needs attention. This paper will include results, comments to results and the essential code.

I've chosen to make all experiments and programs in "Python". My codes can be run an executed online.

## Birthday Experiment
#### Description 
Write a program that takes an integer N from the command line and generate a random sequence of integers between 0 and N – 1. Run experiments to validate the hypothesis that the number of integers generated before the first repeated value is found is ~√(pi)N/2.

#### Execute Program Online

jdoodle.com/a/nln

#### Results & comments:

Results to this problem is as follows:
```
With n=30 runs=1000.0: got result=7.57 hypothesis=6.86468424648
Error diff: -0.0931724905577
With n=365 runs=1000.0: got result=24.553 hypothesis=23.9445329727
Error diff: -0.0247817793065
```
I calculate how much they differs by expected value - actual value / actual value to get percentage error.

This seems very close, the difference is: ~ < 10% error and and ~ < 4% for N=365 and might indicate that the hypothesis ~√(pi)N/2 is true or very close to the actual formula.

The big O notation is the formula. O(√(pi)N/2) But worst case scenario, it could also be O(N). It could potentially, go through the whole available value space before it finds a duplicate. Which is N. 

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

Running the experiments:
```python
n=30
runs=1000.0	
exp1 = experiment(n,runs)
expected = math.sqrt((math.pi*n)/2)


print("With n="+str(n)+" runs="+str(runs)+": got result="+str(exp1)+" hypothesis="+str(expected))
print((expected-exp1)/exp1)

n=365
runs=1000.0	
exp2 = experiment(n,runs) 
expected = math.sqrt((math.pi*n)/2)

print("With n="+str(n)+" runs="+str(runs)+": got result="+str(exp2)+" hypothesis="+str(expected))
print((expected-exp2)/exp2)

```


## Coupon collector problem.
#### Description
Generating random integers as in the previous exercise run experiments to validate the hypothesis that the number of integers generated before all possible values are generated is ~N HN.
#### Execute Program Online

jdoodle.com/a/nBk

#### Results & comments

By running the experiment I would like to conclude that the hypothesis that the formula N*H, where H stands for harmonic number is a quite accurate formula.


I calculate how much they differs by expected value - actual value / actual value to get percentage error.

Those differenced are -0.00154 and -0.018564 which gives us a percentage error of less than 4% error.

```
With n=10 runs=1000.0 result=29.335 expected result=29.289682579
Error diff: -0.00154482430694
With n=100 runs=100.0 result=528.55 expected result=518.737751764
Error diff: -0.0185644654925

```

#### Essential Code

The function to run the coupon problem:
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

Code to print out results compared to expected result:

``` 
n=10
runs=1000.0	
exp1 = experiment(n,runs)
expected =  n*H(n)
print("With n="+str(n)+" runs="+str(runs)+" result="+str(exp1)+" expected result="+str(expected))
print((expected-exp1)/exp1)
n=100
runs=100.0	
exp2 = experiment(n,runs) 
expected = n* H(n)
print("With n="+str(n)+" runs="+str(runs)+" result="+str(exp2)+" expected result="+str(expected))
print((expected-exp2)/exp2)
``` 

I used this formula to calculate the harmonic numbers. The code was taken from [StackOwerFlow](https://stackoverflow.com/questions/404346/python-program-to-calculate-harmonic-series):
 
 ``` 
 def H(n):
    """Returns an approximate value of n-th harmonic number.

       http://en.wikipedia.org/wiki/Harmonic_number
    """
    # Euler-Mascheroni constant
    gamma = 0.57721566490153286060651209008240243104215933593992
    return gamma + log(n) + 0.5/n - 1./(12*n**2) + 1./(120*n**4)



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
This exercise is very open for interpretation and open for many kinds of solutions. 
When translating the idea from physical cards to a program, this is where you can pick easy solutions or more accurate solutions. 
But in essence this problem can be solved by many ways. I've chosen this following way:

- 1. Start at the first Card from the left
- 2. Pick up and look at Card at this first position. 
- 3. Compare all the Cards to the right of the position and find the smallest one. 
- 4. Iterate to the next position and go to step 3. until all Cards are sorted.


This is implemented in selection sort. 

I was choosing Selection sort because we only need to look at two cards at the time. However we need to keep a note or remember the value and position the smallest card.

I am assuming here that we have good memory or a pen to write it down if we need to. 

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
