(ns advent.day-6
  (:require
    [clojure.string :as s]
    [clojure.math.numeric-tower :as math]
    [advent.utils :as utils]))

(def data (utils/get-input-lines "day_6"))

(def items
  (map-indexed
    (fn [i v]
      {:id i
       :pos (mapv read-string (s/split v #", "))})
    data))

(def all-spots
  (for [x (range 53 353)
        y (range 42 353)]
    [x y]))

(defn dist [a b]
  (reduce + (map math/abs [(- (second a) (second b))
                           (- (first a) (first b))])))

(defn dists [pos]
  (map (fn [item] (assoc item :dist (dist pos (:pos item)))) items))

(defn closest-item-id [pos]
  (let [all-dists (dists pos)
        [a b] (take 2 (sort-by :dist all-dists))]
    (if (= (:dist a) (:dist b))
      nil
      (:id a))))

;; (def mapped-items (map closest-item-id all-spots))
;; (def coo-items (zipmap all-spots mapped-items))
;;
;; (remove nil? (set
;;   (map
;;     coo-items
;;     (concat
;;       (for [x [53 353]
;;             y (range 42 353)]
;;         [x y])
;;       (for [x (range 53 353)
;;             y [42 353]]
;;         [x y])))))

(defn run []
  (loop [itms (mapv :pos items)]
    (let [item (first itms)
          all  (vec (rest itms))
          d    (apply + (map (partial dist item) all))]
      (if (< d 10000)
        [item d]
        (do (println (count itms))
            (recur (conj all item)))))))


(def item-pos (mapv :pos items))

(def less-10k
  (map
    (fn [pos]
      (< (apply + (map (partial dist pos) item-pos)) 10000))
    all-spots))
