(ns day-9.code)

(def test-data
  (clojure.string/split-lines (slurp "day_9/test.txt")))

(def final-data
  (clojure.string/split-lines (slurp "day_9/input.txt")))

(defn build-board [lines]
  (let [board (->> lines
                   (mapv #(->> (clojure.string/split % #"")
                               (mapv (fn [e] (Integer/parseInt e))))))]
    {:board board
     :y (count board)
     :x (count (first board))}))

(defn lookup [board x y]
  (-> board :board (nth y) (nth x)))

(defn valid-coord [board x y]
  (and (>= x 0)
       (>= y 0)
       (< x (:x board))
       (< y (:y board))))

(defn surround [board x y]
  (filter
    #(apply valid-coord board %)
    [[(dec x) y]
     [(inc x) y]
     [x (inc y)]
     [x (dec y)]]))

(defn find-lows [board]
  (reduce
    (fn [coll y]
      (reduce
        (fn [coll x]
          (let [match (lookup board x y)
                small (apply min (map #(apply lookup board %)
                                      (surround board x y)))]
            (if (< match small)
              (conj coll [match x y])
              coll)))
        coll
        (range (:x board))))
    []
    (range (:y board))))

(def part-1
  (apply + (map inc (map first (find-lows (build-board final-data))))))

(defn basin-size [board ix iy]
  (-> (loop [visited #{}
             check [[ix iy]]
             match []]
        (if-let [[x y] (first check)]
          (let [v (lookup board x y)
                adjs (->> (surround board x y)
                          (remove visited)
                          (filter #(let [nv (apply lookup board %)]
                                     (and (< nv 9) (> nv v)))))]
            (recur (conj visited [x y])
                   (concat (rest check) adjs)
                   (conj match [x y])))
          match))
      distinct
      count))

(defn basin-score [board]
  (->> (find-lows board)
       (map rest)
       (map #(apply basin-size board %))
       (sort)
       reverse
       (take 3)
       (apply *)))

(def part-2
  (basin-score (build-board final-data)))
