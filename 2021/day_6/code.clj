(ns day-6.code)

(defn read-input [input]
  (read-string (str "[" input "]")))

(def test-data
  (read-input "3,4,3,1,2"))

(def final-data
  (read-input (slurp "day_6/input.txt")))

(defn brute-progress [fish]
  (mapcat
    #(cond
       (zero? %)
         [6 8]
       :else
         [(dec %)])
    fish))

(defn do-progress [progress-fn fish days]
  (reduce (fn [v _]
            (progress-fn v))
          fish
          (range days)))

(defn progress-indexed
  "Progress with indexed fish. Keep counts of fish per day, instead
  of a flat list.

  [1 2 3 0 0 0 1 3 0]

  becomes

  days:
   0           6   8
   |           |   |
  [2 3 0 0 0 2 3 0 2]"
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
  (count (do-progress brute-progress final-data 80)))

(def part-2
  (reduce + (do-progress progress-indexed (index-fish (frequencies final-data)) 256)))
