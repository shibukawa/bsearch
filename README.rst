Binary Search for golang
=============================

It provides similar functionality to ``sort.Search()`` of standard golang library.

It searches element or insert position from sorted array.

bsearch.Search(n int, compare CompareFunc) int
-----------------------------------------------

It returns index of the target element. It returns -1 if the element is not in the list.

.. code-block:: go

   items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
   target := 12
   result := Search(len(items), func(i int) CompareResult {
       if target > items[i] {
           return Smaller
       } else if target == items[i] {
           return Equal
       } else {
           return Bigger
       }
   })
   // result -> -1

bsearch.FindInsertPosition(n int, compare CompareFunc) int
-------------------------------------------------------------

It is almost as same as ``sort.Search()``. It returns insert poisition for the target element in the list.

.. code-block:: go

   items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
   target := 12
   result := Search(len(items), func(i int) CompareResult {
       if target > items[i] {
           return Smaller
       } else if target == items[i] {
           return Equal
       } else {
           return Bigger
       }
   })
   // result -> 6

Constants
---------

The following constants are used from callback functions:

.. code-block:: go

   type CompareResult int

   const (
       Smaller CompareResult = -1
       Equal CompareResult = 0
       Bigger CompareResult = 1
   )

The following type is the type of callback functions:

.. code-block:: go

   type CompareFunc func(int)CompareResult

License
---------

MIT
