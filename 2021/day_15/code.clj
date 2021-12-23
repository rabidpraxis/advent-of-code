(ns day-15.code
  (:require
    [clojure.data.priority-map :refer [priority-map-keyfn]]
    [clojure.string :as str]))

(def test-data
  (str/split-lines (slurp "day_15/test.txt")))

(def final-data
  (str/split-lines (slurp "day_15/input.txt")))

(defn build-board [lines]
  (let [grid (mapv #(mapv read-string (str/split % #"")) lines)]
    {:grid grid
     :y (count grid)
     :x (count (first grid))}))

(defn lookup [board x y]
  (-> board :grid (nth y) (nth x)))

(defn valid-coord
  "filter out overflow coordinates"
  [board x y]
  (and (>= x 0) (>= y 0) (< x (:x board)) (< y (:y board))))

(defn neighbors [board x y]
  (filter
    #(apply valid-coord board %)
    [[(dec x) y] [(inc x) y] [x (inc y)] [x (dec y)]]))

(defn build-graph [board]
  (let [{:keys [x y]} board
        coords (for [nx (range x) ny (range y)] [nx ny])]
    (reduce
      (fn [coll coord]
        (assoc coll coord (apply neighbors board coord)))
      {}
      coords)))

(defn dist [board _ coord]
  (apply lookup board coord))

(defn ^:private generate-route [node came-from]
  (loop [route '()
         node node]
    (if (came-from node)
      (recur (cons node route) (came-from node))
      route)))

(defn route
  "Extracted from https://github.com/arttuka/astar. For learnin"
  [graph dist start goal]
  (loop [visited {}
         queue (priority-map-keyfn first start [0 0 nil])]
    (when (seq queue)
      (let [[current [_ current-score previous]] (peek queue)
            visited (assoc visited current previous)]
        (if (= current goal)
          (generate-route goal visited)
          (recur visited (reduce (fn [queue node]
                                   (let [score (+ current-score (dist current node))]
                                     (if (and (not (contains? visited node))
                                              (or (not (contains? queue node))
                                                  (< score (get-in queue [node 1]))))
                                       (assoc queue node [score score current])
                                       queue)))
                                 (pop queue)
                                 (graph current))))))))

(defn optimal-path [board]
  (let [{:keys [x y]} board
        graph (build-graph board)]
    (route graph (partial dist board) [0 0] [(dec x) (dec y)])))

(defn expanded-board [lines m]
  (let [grid (mapv #(mapv read-string (str/split % #"")) lines)
        x (count grid)
        y (count (first grid))
        xl (* m x)
        yl (* m y)
        final-grid (for [nx (range xl)]
                     (for [ny (range yl)]
                       (let [mx (mod nx x)
                             xmult (int (/ nx x))
                             my (mod ny y)
                             ymult (int (/ ny y))
                             sval (-> grid (nth mx) (nth my))
                             val (+ (+ xmult ymult) sval)]
                         (if (> val 9)
                           (rem val 9)
                           val))))]
    {:grid (into [] (map #(into [] %) final-grid)) :y xl :x yl}))

(defn sum-path [board]
  (->> board
       optimal-path
       (map #(apply lookup board %))
       (apply +)))

(def part-1
  (sum-path (build-board final-data)))

(def part-2
  (sum-path (expanded-board final-data 5)))

(comment
  (map #(apply str %) (:grid (expanded-board test-data 5))))
