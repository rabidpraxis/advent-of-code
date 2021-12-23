(ns day-17.code
  (:require
    [clojure.set :refer [difference]]
    [clojure.string :as str]))

(def final-board
  {:x [150 193]
   :y [-136 -86]})

(def test-board
  {:x [20 30]
   :y [-10 -5]})

(defn step [{start :start
             [vx vy] :vel
             [px py] :pos
             max-y :max-y
             :or {max-y 0}}]
  (let [y (+ py vy)]
    {:vel [(cond
             (= vx 0) 0
             (> vx 1) (- vx 1)
             (< vx 1) (+ vx 1)
             :else 0)
           (- vy 1)]
     :start start
     :max-y (if (> y max-y)
              y
              max-y)
     :pos [(+ px vx) y]}))

(defn go [[x y] board]
  (->> {:pos [0 0]
        :start [x y]
        :vel [x y]}
       (iterate step)
       (filter
         (fn [{[px py] :pos}]
           (let [{[x1 x2] :x
                  [y1 y2] :y} board]
             (or (and (<= x1 px x2)
                      (<= y1 py y2))
                 (or (> px x2)
                     (< py y1))))))
       first))

(defn in-range? [{[px py] :pos} {[x1 x2] :x [y1 y2] :y}]
  (and (<= x1 px x2) (<= y1 py y2)))


(defn part-1
  (apply max
         (for [xvel (range 200)
               yvel (range 200)
               :let [res (go [xvel yvel] final-board)]
               :when (in-range? res final-board)]
           (:max-y res))))

(defn part-2
  (count (for [xvel (range -150 250)
               yvel (range -150 250)
               :let [res (go [xvel yvel] final-board)]
               :when (in-range? res final-board)]
           (:start res))))

(comment
  (def matches
    (->> (str/split
           "23,-10  25,-9   27,-5   29,-6   22,-6   21,-7   9,0     27,-7   24,-5 25,-7   26,-6   25,-5   6,8     11,-2   20,-5   29,-10  6,3     28,-7 8,0     30,-6   29,-8   20,-10  6,7     6,4     6,1     14,-4   21,-6 26,-10  7,-1    7,7     8,-1    21,-9   6,2     20,-7   30,-10  14,-3 20,-8   13,-2   7,3     28,-8   29,-9   15,-3   22,-5   26,-8   25,-8 25,-6   15,-4   9,-2    15,-2   12,-2   28,-9   12,-3   24,-6   23,-7 25,-10  7,8     11,-3   26,-7   7,1     23,-9   6,0     22,-10  27,-6 8,1     22,-8   13,-4   7,6     28,-6   11,-4   12,-4   26,-9   7,4 24,-10  23,-8   30,-8   7,0     9,-1    10,-1   26,-5   22,-9   6,5 7,5     23,-6   28,-10  10,-2   11,-1   20,-9   14,-2   29,-7   13,-3 23,-5   24,-8   27,-9   30,-7   28,-5   21,-10  7,9     6,6     21,-5 27,-10  7,2     30,-9   21,-8   22,-7   24,-9   20,-6   6,9     29,-5 8,-2    27,-8   30,-5   24,-7"
           #"\s+")
         (map #(mapv read-string (str/split % #",")))
         set))

  (in-range? (go [7 -1] test-board) test-board)
  (difference
    matches
    (set (count (for [xvel (range -55 55)
               yvel (range -55 55)
               :let [res (go [xvel yvel] test-board)]
               :when (in-range? res test-board)]
           (:start res))))))
