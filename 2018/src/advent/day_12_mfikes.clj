(ns advent.day-12-mfikes
  (:require
   [advent.utils :as utils]
   [clojure.string :as string]))

;; (def input (-> "advent_2018/day_12/input" io/resource slurp))
(def input (utils/get-input-lines "day_12"))

(defn parse-note [note]
  (when (= "#" (subs note 9))
    (map #{\#} (subs note 0 5))))

(defn parse-state [state]
  (into (sorted-set) (keep-indexed (fn [idx pot]
                            (when (= \# pot) idx))
              state)))

(second (parse-input input))

(defn parse-input [input]
  (let [[initial-state-line _ & notes] input]
    [(parse-state (subs initial-state-line 15))
     (into #{} (keep parse-note notes))]))

(defn generation [notes state]
  (reduce (fn [state' idxs]
            (if (notes (map #(when (state %) \#) idxs))
              (conj state' (nth idxs 2))
              state'))
    (sorted-set)
    (partition 5 1
      (range (- (first state) 5) (inc (+ (first (rseq state)) 5))))))

(defn states [input]
  (let [[state notes] (parse-input input)]
    (iterate (partial generation notes) state)))

(defn solve [states n]
  (reduce + (nth states n)))

(defn solve-extrapolating [states n]
  (let [shifted? (fn [state-2 state-1]
                   (apply = (map - state-1 state-2)))
        idx      (first (keep-indexed (fn [idx [s1 s2]]
                                        (when (shifted? s1 s2)
                                          idx))
                          (map vector states (rest states))))
        diff     (- (solve states (inc idx)) (solve states idx))]
    (+ (solve states idx) (* diff (- n idx)))))

(defn part-1 []
  (solve (states input) 20))

(defn part-2 []
  (solve-extrapolating (states input) 50000000000))
