(ns day-11.code
  (:require
    [clojure.string :as str]))

(def test-data
  (str/split-lines (slurp "day_11/test.txt")))

(def final-data
  (str/split-lines (slurp "day_11/input.txt")))

(defn build-board [lines]
  (let [board (mapv #(mapv read-string (str/split % #"")) lines)]
    {:board board
     :step 0
     :flashct 0
     :y (count board)
     :x (count (first board))}))

(defn lookup [board x y]
  (get-in board [:board y x]))

(defn valid-coord
  "filter out overflow coordinates"
  [board x y]
  (and (>= x 0)
       (>= y 0)
       (< x (:x board))
       (< y (:y board))))

(defn neighbors
  "find valid neighboring coords (including diags)"
  [board x y]
  (for [nx (range (- x 1) (+ x 2))
        ny (range (- y 1) (+ y 2))
        :when (and (not= [nx ny] [x y])
                   (valid-coord board nx ny))]
    [nx ny]))

(defn inc-board [board]
  (update board :board
    (fn [b]
      (mapv #(mapv inc %) b))))

(defn inc-pos
  "increment a specific position and return updated board with
  incremented value"
  [board x y]
  (let [v (lookup board x y)
        nv (inc v)]
    [(assoc-in board [:board y x] nv) nv]))

(defn find-flashes
  [board]
  (for [y (range (:y board))
        x (range (:x board))
        :when (>= (lookup board x y) 10) ]
    [x y]))

(defn step-flashes
  "increment all items and return a new board and all neighbors
  of currently flashing numbers"
  [board]
  (let [step-board (inc-board board)
        flashes (find-flashes step-board)]
    [step-board (mapcat
                  #(apply neighbors step-board %)
                  flashes)]))

(defn process-flashes
  "given a set of coordinates to increment, bump and check for
  new flashes. If there are new flashes from the resulting bump,
  then add to the list of flashes and keep incrementing."
  [board flashes]
  (loop [flashes flashes
         board board]
    (if-let [[x y] (first flashes)]
      (let [[board nv] (inc-pos board x y)
            nflashes (if (= 10 nv)
                       (concat (rest flashes) (neighbors board x y))
                       (rest flashes))]
        (recur nflashes board))
      board)))

(defn zero-board-pos [board x y]
  (assoc-in board [:board y x] 0))

(defn zero-flashes
  "zero out any previously flashing items and update the
  flash counts"
  [board]
  (let [flashes (find-flashes board)]
    (-> (reduce #(apply zero-board-pos %1 %2) board flashes)
        (update :flashct (partial + (count flashes))))))

(defn step [board]
  (->> (update board :step inc)
       step-flashes
       (apply process-flashes)
       zero-flashes))

(def part-1
  (-> (iterate step (build-board final-data))
      (nth 100)
      :flashct))

(defn all-flashing? [board]
  (every? #(every? zero? %) (:board board)))

(def part-2
  (->> (iterate step (build-board final-data))
       (filter all-flashing?)
       first
       :step))
