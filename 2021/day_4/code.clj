(ns day-4.code
  (:require
    [clojure.string :as str]))

(def final-data
  (str/split-lines (slurp "day_4/input.txt")))

(def test-board-data
  "22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7")

(defn parse-scores [scores]
  (map #(Integer/parseInt %) (str/split scores #",")))

(defn split-board [lines]
  (map #(-> % (str/trim) (str/split #"\s+")) lines))

(defn index [item]
  (map vector (range (count item)) item))

(defn build-board [board]
  (reduce
    (fn [coll [x row]]
      (reduce
        (fn [coll [y id]]
          (assoc coll (Integer/parseInt id) [x y false]))
        coll
        (index row)))
    {}
    (index board)))

(defn boards [lines]
  (->> lines
       (partition 5 6)
       (map (comp build-board split-board))))

(defn winning-board? [board]
  (or (->> (partition 5 (sort-by first (vals board)))
           (map #(map last %))
           (map #(every? true? %))
           (some true?))
      (->> (partition 5 (sort-by second (vals board)))
           (map #(map last %))
           (map #(every? true? %))
           (some true?))))

(defn apply-score [board score]
  (if (contains? board score)
    (update board score
      (fn [[x y _]]
        [x y true]))
    board))

(def test-scores
  (parse-scores "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"))

(def test-boards
  (boards (str/split-lines test-board-data)))

(defn find-first-winner [scores boards]
  (let [score (first scores)
        updated-boards (map #(apply-score % score) boards)]
    (if-let [winner (first (filter winning-board? updated-boards))]
      [score winner]
      (recur (rest scores) updated-boards))))

(defn unmarked-score [score board]
  (let [unmarked (filter (comp false? last last) (seq board))
        unmarked-tot (reduce #(+ %1 (first %2)) 0 unmarked)]
    (* score unmarked-tot)))

(def final-scores
  (parse-scores (first final-data)))

(def final-boards
  (boards (nthrest final-data 2)))

(def part-1
  (apply unmarked-score (find-first-winner final-scores final-boards)))

(defn find-last-winner [scores boards losers]
  (let [score (first scores)
        updated-boards (map #(apply-score % score) boards)]
    (if (every? winning-board? updated-boards)
      [score (apply-score (last losers) score)]
      (recur (rest scores) updated-boards (remove winning-board? updated-boards)))))

(def part-2
  (apply unmarked-score (find-last-winner final-scores final-boards nil)))

(comment
  (apply unmarked-score (find-first-winner test-scores test-boards))
  (apply unmarked-score (find-last-winner test-scores test-boards nil)))
