(ns advent.day-1
  (:require
    [clojure.string :as string]
    [advent.answers :as answers]
    [advent.utils :as utils]))

(def data (utils/get-input-lines "day_1"))

(def signed-data (map read-string data))

(defn find-duplicate [xs]
  (reduce
    (fn [history freq]
      (if (history freq)
        (reduced freq)
        (conj history freq)))
    #{} xs))

(defn part-1 []
  (apply + signed-data))

(defn part-1? []
  (= answers/day-1-1 (part-1)))

(defn part-2 []
  (find-duplicate (cons 0 (reductions + (cycle signed-data)))))

(defn part-2? []
  (= answers/day-1-2 (part-2)))
