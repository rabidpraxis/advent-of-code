(ns advent.day-2
  (:require
    [clojure.string :as string]
    [clojure.set :as set]
    [advent.answers :as answers]
    [advent.utils :as utils]))

(def data (utils/get-input-lines "day_2"))

(defn get-counts [ids ct]
  (->> ids
    (map (comp set vals frequencies))
    (filter #(% ct))
    count))

(defn part-1 []
  (* (get-counts data 2) (get-counts data 3)))

(defn part-1? []
  (= answers/day-2-1 (part-1)))

;; Part 2
(defn transpose [& xs]
  (apply map list xs))

(def target-ct 25)

(defn similar [a b]
  (->> (transpose a b)
       (filter #(apply = %))
       (map first)
       (string/join "")))

(defn part-2 []
  (loop [items data]
    (let [item (first items)
          tail (rest items)]
      (if-let [found (first (filter #(= (count %) target-ct)
                                    (map (partial similar item) tail)))]
        found
        (recur tail)))))

(defn part-2? []
  (= answers/day-2-2 (part-2)))
