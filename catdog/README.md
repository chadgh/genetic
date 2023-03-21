# Cat and Dog Care Problem
Optimizing for earnings for Carlos and Clarita's new pet sitting business with a genetic algorithm.

Note: This solution assumes that Carlos can help a little with the pampering time.

## The Problem
Carlos and Clarita have successfully found a way to represent all of the combinations of cats and
dogs that they can board based on all of the following constraints.

* Space: Cat pens require of space, while dog runs require . Carlos and Clarita have
up to available in the storage shed for pens and runs, while still leaving enough room
to move around the cages.
* Feeding Time: Carlos and Clarita estimate that cats will require minutes twice a
day—morning and evening—to feed and clean their litter boxes, for a total of minutes per
day for each cat. Dogs will require minutes twice a day to feed and walk, for a total of
minutes per day for each dog. Carlos can spend up to hours each day for the morning and
evening feedings, but needs the middle of the day off for baseball practice and games.
* Pampering Time: The twins plan to spend minutes each day brushing and petting each cat,
and minutes each day bathing or playing with each dog. Clarita needs time off in the
morning for swim team and evening for her art class, but she can spend up to hours during
the middle of the day to pamper and play with the pets.
* Start-up Costs: Carlos and Clarita plan to invest much of the they earned from their
last business venture to purchase cat pens and dog runs. It will cost for each cat pen and
for each dog run.

Now they are trying to determine how many of each type of pet they should plan to accommodate.
Of course, Carlos and Clarita want to make as much money as possible from their business, so they
need to pay attention to both their daily income as well as their daily costs. They plan to charge
per day for boarding each cat and per day for each dog. They estimate that each cat will require
per day in food and supplies, and that each dog will require per day in costs.

After surveying the community regarding the pet boarding needs, Carlos and Clarita are confident
that they can keep all of their boarding spaces filled for the summer.

So the problem is, how many of each type of pet should they prepare for in order to make as
much money as possible?

What combination of cats and dogs do you think will make the most money? What
recommendations would you give to Carlos and Clarita, and what argument would you use to
convince them that your recommendation is reasonable? 


## The Solution
The fitness.go file shows the constraint numbers.

The three solutions it comes up with are either:
* 10 cats, 12 dogs
* 15 cats, 10 dogs
* 20 cats, 8 dogs

Which all produce a total earnings of $240 per day.

It also finds 25 cats, 6 dogs as a solution with the same daily
earnings, but this solution assumes that pampering and feeding
times are pulled from the same 16 hour pool.