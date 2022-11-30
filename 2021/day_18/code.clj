(ns day-18.code
  (:require
    [clojure.string :as str]
    [clojure.walk :refer [postwalk]]
    [clojure.math.combinatorics :refer [combinations]]
    [hyperfiddle.rcf :refer [tests ! %]]))

(hyperfiddle.rcf/enable!)

(def test-data
  (str/split-lines (slurp "day_18/test.txt")))

(def final-data
  (str/split-lines (slurp "day_18/input.txt")))

(defn r [s]
  (read-string s))

(defn gather-node-pos
  "list all node positions"
  ([v]
   (gather-node-pos v [] []))
  ([v path found]
   (reduce
     (fn [found [idx v]]
       (if (vector? v)
         (gather-node-pos v (conj path idx) found)
         (conj found (conj path idx))))
     found
     (map-indexed list v))))

(defn *find
  ([coll check-fn]
   (:found (*find check-fn coll [])))
  ([check-fn coll path]
   (reduce
     (fn [path [idx v]]
       (if (:found path)
         path
         (let [npath (conj path idx)]
           (cond
             (check-fn v npath)
               {:found npath}
             (vector? v)
               (let [ret (*find check-fn v npath)]
                 (if (:found ret)
                   ret
                   (pop ret)))
             :else
               path))))
     path
     (map-indexed list coll))))

(defn find-explosion
  [coll]
  (*find coll (fn [v p] (and (vector? v)
                             (int? (first v))
                             (int? (second v))
                             (>= (count p) 4)))))

(defn find-split
  [coll]
  (*find coll (fn [v _] (and (int? v) (> v 9)))))

(defn explode [p]
  (if-let [u (find-explosion p)]
    (let [[av bv] (get-in p u)
          pos (gather-node-pos p)
          bp (->> pos
                  (partition 2 1)
                  (filter (fn [[_ v]] (= v (conj u 0))))
                  ffirst)
          ap (->> pos
                  reverse
                  (partition 2 1)
                  (filter (fn [[_ v]] (= v (conj u 1))))
                  ffirst)]
      (cond-> p
        true (assoc-in u 0)
        bp (update-in bp #(+ % av))
        ap (update-in ap #(+ % bv))))
    p))

(defn split [p]
  (if-let [u (find-split p)]
    (update-in p u (fn [e]
                     [(int (Math/floor (/ e 2)))
                      (int (Math/ceil (/ e 2)))]))
    p))

(defn sum [v]
  (postwalk #(if (vector? %) (+ (* 3 (first %)) (* 2 (second %))) %) v))

(defn apply-change
  "apply f against p until (not= p (f p))"
  [p f]
  (loop [p p]
    (let [res (-> p f)]
      (if (= res p) p (recur res)))))

(defn add-change
  [a b]
  (apply-change [a b] #(-> % (apply-change explode) split)))

(defn factor [lines]
  (let [d (map r lines)]
    (reduce add-change (first d) (rest d))))

(tests
  "explodes a single value"
  (explode (r "[[[[[9,8],1],2],3],4]"))
  := (r "[[[[0,9],2],3],4]")

  "explodes missing seperate val"
  (explode (r "[7,[6,[5,[4,[3,2]]]]]"))
  := (r "[7,[6,[5,[7,0]]]]")

  "explodes inside another array"
  (explode (r "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"))
  := (r "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")

  "explodes endings"
  (explode (r "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"))
  := (r "[[3,[2,[8,0]]],[9,[5,[7,0]]]]")

  "Basic change"
  (add-change [[[[4,3],4],4],[7,[[8,4],9]]] [1,1])
  := [[[[0,7],4],[[7,8],[6,0]]],[8,1]]

  "Change 2"
  (add-change
    [[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
    [7,[[[3,7],[4,3]],[[6,3],[8,8]]]])
  := [[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]

  "sums"
  (sum (r "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")) := 1384)

(def part-1
  (sum (factor final-data)))

(def part-2
  (apply max (mapcat
               (fn [[a b]]
                 [(sum (add-change a b))
                  (sum (add-change b a))])
               (combinations (map r final-data) 2))))
