(ns day-6.code)

(defn read-input [input]
  (read-string (str "[" input "]")))

(def test-data
  (read-input "3,4,3,1,2"))

(def final-data
  (read-input (slurp "day_6/input.txt")))

(defn brute-progress [fish]
  (mapcat #(if (zero? %)
             [6 8]
             [(dec %)])
          fish))

(defn repeat-fn
  "reduce over value (d) by specific count (ct) with function (f)"
  [f ct d]
  (reduce (fn [v _] (f v)) d (range ct)))

(defn indexed-progress
  "Find progress by indexing fish counts per day, vs
  building up a flat list. For example, counts for each
  9 days would make an indexed list like:

  [1 2 3 0 0 0 1 3 0]

  which after another day becomes:

  days:
   0           6   8
   |           |   |
  [2 3 0 0 0 1 4 0 1]"
  [fish]
  (let [today (first fish)
        field (vec (rest fish))]
    (-> field
        (update 6 + today)
        (conj today))))

(defn index-fish [freq]
  (reduce
    (fn [v d]
      (assoc v d (or (get freq d) 0)))
    []
    (range 9)))

(def part-1
  (count (repeat-fn brute-progress 80 final-data)))

(def part-2
  (->>
    (frequencies final-data)
    index-fish
    (repeat-fn indexed-progress 256)
    (reduce +)))
