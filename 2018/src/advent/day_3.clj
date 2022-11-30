(ns advent.day-3
  (:require
    [clojure.string :as string]
    [clojure.set :as set]
    [advent.answers :as answers]
    [advent.utils :as utils]))

(def data (utils/get-input-lines "day_3"))

(defn extract-data [v]
  (let [re #"#([0-9]+) @ (\d+),(\d+): (\d+)x(\d+)"
        [id x y w h] (map read-string (rest (re-matches re v)))]
    {:id id :x x :y y :w w :h h}))

(defn inject-coords [coll item]
  (let [{:keys [id x y w h]} item
        coords (for [xr (range x (+ x w)) yr (range y (+ y h))] [xr yr])]
    (reduce #(update %1 %2 conj id) coll coords)))

(def final-coords (reduce inject-coords {} (map extract-data data)))

(defn part-1 []
  (count (filter #(> (count %) 1) (vals final-coords))))

(defn part-1? []
  (= answers/day-3-1 (part-1)))

;; Part 2
(defn part-2 []
  (let [all-vals   (vals final-coords)
        mult-set   (set (flatten (filter #(> (count %) 1) all-vals)))
        single-set (set (flatten (filter #(= (count %) 1) all-vals)))]
    (first (set/difference single-set mult-set))))

(defn part-2? []
  (= answers/day-3-2 (part-2)))
