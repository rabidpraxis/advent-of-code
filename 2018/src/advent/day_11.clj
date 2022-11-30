(ns advent.day-11
  (:require
    [clojure.set :as set]
    [clojure.string :as string]
    [clojure.core.matrix :as matrix]
    [advent.answers :as answers]
    [advent.utils :as utils]))

(matrix/set-current-implementation :vectorz)

(defn hundreth [num]
  (mod (quot num 100) 10))

(defn get-power [serial-num [x y]]
  (let [rack-id   (+ x 10)
        power-lvl (* rack-id y)
        w-serial  (+ power-lvl serial-num)
        rack-mult (* w-serial rack-id)
        h-dig     (hundreth rack-mult)]
    (- h-dig 5)))

(defn grid [serial]
  (vec (for [x (range 1 301)]
    (vec (for [y (range 1 301)]
      (get-power serial [x y]))))))

(def *grid (matrix/array (grid 7857)))

(defn get-square-sum [x y size]
  (matrix/esum (matrix/select *grid (range x (+ x size)) (range y (+ y size)))))

(defn all-powers [size]
  (for [x (range (- 300 size))
        y (range (- 300 size))]
    (let [ret   [(inc x) (inc y) size]
          power (get-square-sum x y size)]
      [ret power])))

(defn max-power [powers]
  (apply max-key second powers))

(defn part-1 []
  (let [[[x y _] _] (max-power (all-powers 3))]
    (string/join "," [x y])))

(defn part-2 []
  (let [all (pmap
              (comp max-power all-powers)
              (range 1 300))
        [[x y size] _] (max-power all)]
    (string/join "," [x y size])))
