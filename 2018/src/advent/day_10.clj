(ns advent.day-10
  (:require
    [clojure.string :as string]
    [advent.answers :as answers]
    [advent.utils :as utils]))

(def data (utils/get-input-lines "day_10_test"))
(def data (utils/get-input-lines "day_10"))

(defn extract-coords [s]
  (let [[pos vel] (map (fn [s]
                         (map read-string (string/split (second s) #",")))
                       (re-seq #"<(.*?)>" s))]
    {:pos pos :vel vel}))

(def coords (map extract-coords data))

(defn move [{[px py] :pos [vx vy] :vel}]
  {:pos [(+ px vx) (+ py vy)]
   :vel [vx vy]})

(defn move-and-count [{:keys [coords iteration]}]
  (let [coords (map move coords)
        counts (->> (map (comp first :pos) coords)
                 frequencies
                 (map second))]
    {:coords coords
     :iteration (inc iteration)
     :count-avg (float (/ (reduce + counts) (count counts)))}))

(def possible
  (->> {:coords coords :iteration 0 :count-avg 0}
       (iterate move-and-count)
       (take 20000)
       (apply max-key :count-avg)))

(defn part-1 []
  (let [poss (map :pos (:coords possible))
        xset (map first poss)
        yset (map second poss)
        minx (apply min xset)
        maxx (apply max xset)
        miny (apply min yset)
        maxy (apply max yset)
        normal-pos (set (map (fn [[x y]] [(- x minx) (- y miny)]) poss))
        pieces (for [y (range 0 (+ 1 (- maxy miny)))]
                 (for [x (range 0 (+ 1 (- maxx minx)))]
                   (if (normal-pos [x y]) \# \.)))]
    (println (string/join "\n" (map (partial string/join "")  pieces)))))

(defn part-2 []
  (:iteration possible))
