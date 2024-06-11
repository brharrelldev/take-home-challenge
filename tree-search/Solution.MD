## Solution

**Disclaimer: Put the code in Tree-Search package**

I looked up a few different types of trees.  I always assumed trees could not have duplicates, so I wasn't quite sure about any tree that could support them. Here are a few

B-Tree - This is usually used for indexing by creating shallower trees. However because of this, its unlikley to be able to support duplicate keys.  At least if implemented correctly

AVL-Tree - The insert requires rotation.  This can be very expensive, so in this case it wouldn't make sense to support insertion.  Like B-Trees it is also balanced, meaning that by nature it is a lot more shallow

Red-Black-Tree - Due to the coloring algorithm, duplicates can't really exist here.  It is also balanced, so this means the tree is likely to be shallow


The only tree that seems to support duplicates is a Binary Search Tree.

In this solution I decided to keep count at the time of insertion.  I tried a few other approaches like counting after each traversal.  Only way I could find is to use the count as
a parameters to the CheckDuplicatesID and increment it.  But I guess this is kind of cheating because I just made count apart of the node. 

CheckID is a normal tree traversal, so as a result its O(N). I'm not a math whiz, but that seems to be proof with this equation

𝑇(𝑛)=𝑇(𝑛−1−∣𝑟𝑛𝑜𝑑𝑒∣)+𝑇(𝑛−1−∣ℓ𝑛𝑜𝑑𝑒∣)+1

In this case T being the tree and being the sum of the left and right node +1. So from this equation, it is proof by induction to Big(N)

CheckDuplicateID seems to be an inOrder traversal, but I believe all tree traversals are Big(N)