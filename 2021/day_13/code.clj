(ns day-13.code
  (:require
    [clojure.string :as str]))

(def test-data
  (str/split-lines (slurp "day_13/test.txt")))

(def final-data
  (str/split-lines (slurp "day_13/input.txt")))

(defn board [data]
  {:board (->> (filter #(re-matches #"^\d.*" %) data)
               (mapv #(mapv read-string (str/split % #","))))

   :folds (->> (filter #(re-matches #"^fold.*" %) data)
               (map #(let [[_ dir pos] (re-find #"(x|y)=(\d*)" %)]
                       [dir (read-string pos)])))})

(defn fold [t v [x y]]
  (if (= t "y")
    (if (> y v)
      [x (- v (- y v))]
      [x y])
    (if (> x v)
      [(- v (- x v)) y]
      [x y])))

(def part-1
  (let [{:keys [board folds]} (board final-data)
        [t v] (first folds) ]
    (->> (map (partial fold t v) board)
         distinct
         count)))

(defn render-board [board]
  (let [x (inc (apply max (map first board)))
        y (inc (apply max (map second board)))
        grid (vec (for [_ (range y)]
                    (vec (replicate x " "))))]
    (->> (reduce (fn [grid [x y]] (assoc-in grid [y x] "x")) grid board)
         (map #(apply str %))
         (str/join "\n"))))

(def part-2
  (let [{:keys [board folds]} (board final-data)]
    (->> (reduce
           (fn [board [t v]]
             (map (partial fold t v) board))
           board
           folds)
         render-board
         print)))
