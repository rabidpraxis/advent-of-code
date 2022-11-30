(ns advent.day-15
  (:require
    [clojure.string :as string]
    [advent.answers :as answers]
    [advent.utils :as utils])
  (:import java.lang.Math))

;; Utils
;;
(defn replace-char [input char coord]
  (let [[x y] coord]
    (update input y #(str (subs % 0 x) char (subs % (inc x) (count %))))))

(defn print-input [input]
  (println (string/join \newline input)))

(defn clean-pieces [input pieces]
  (reduce
    (fn [input {:keys [coord]}]
      (replace-char input "." coord))
    input
    pieces))

(defn vec-add [[x1 y1] [x2 y2]]
  [(+ x1 x2) (+ y1 y2)])

(defn dist [a b]
  (reduce + (map #(Math/abs (- %1 %2)) a b)))

;; Start
;;
(def input (utils/get-input-lines "day_15_target_training"))

(def board-config
  {:width (apply max (map count input))
   :height (count input)})

(def piece-map
  {\E :elf
   \G :goblin
   \# :wall
   \. :space})

(def ipiece-map
  (apply hash-map (mapcat reverse piece-map)))

(defn get-piece [piece]
  (:piece piece))

(defn wall? [piece]
  (= (get-piece piece) :wall))

(defn elf? [piece]
  (= (get-piece piece) :elf))

(defn goblin? [piece]
  (= (get-piece piece) :goblin))

(defn character? [piece]
  (or (elf? piece) (goblin? piece)))

(defn character-or-wall? [piece]
  (or (character? piece) (wall? piece)))

(def all-pieces
  (let [piece-ids (atom 0)]
    (for [x (range (:width board-config))
          y (range (:height board-config))]
      (let [raw-piece (piece-map (get (get input y) x))
            piece {:coord [x y]
                   :piece raw-piece}]
        (if (character? piece)
          (assoc piece
                 :id (swap! piece-ids inc)
                 :hp 300)
          piece)))))

(def elves   (partial filter elf?))
(def goblins (partial filter goblin?))

(defn enemies [set piece]
  (let [filterfn (if (elf? piece) goblins elves)]
    (filterfn set)))

(def board-pieces
  (map #(if (character? %)
          (assoc % :piece :space)
          %) all-pieces))

(defn index-pieces [pieces]
  (apply hash-map (mapcat identity (map #(do [(:coord %) %]) pieces))))

(def indexed-board
  (index-pieces board-pieces))

(def characters
  (atom (filter character? all-pieces)))

(def clean-input (clean-pieces input @characters))

(defn with-pieces
  "Merge all pieces into one input"
  [input & piece-group]
  (reduce
    (fn [memo piece]
      (reduce
        (fn [input {:keys [coord piece]}]
          (replace-char input (if (keyword? piece)
                                (ipiece-map piece)
                                piece) coord))
        memo
        piece))
    input
    piece-group))


(defn neighbor-coords [coord]
  (map (partial vec-add coord) [[0 1] [1 0] [-1 0] [0 -1]]))

(defn into-piece [coords piece]
  (map #(do {:piece piece :coord %}) coords))

(defn in-range-coords
  [indexed {:keys [coord]}]
  (->> (neighbor-coords coord)
       (map indexed)
       (filter (complement character-or-wall?))
       (map :coord)))

(defn in-range-pieces [indexed piece]
  (into-piece (in-range-coords indexed piece) \?))

(defn in-range-enemies [index characters character]
  (mapcat (partial in-range-pieces index)
          (enemies characters character)))


(defn get-trails
  "Get all trails from base to target"
  [base target indexed]
  (loop [visited #{}
         trail    []
         frontier [{:score 0 :coord (:coord base)}]
         round    0]
    (if-let [current (first frontier)]
      (let [frontiers (reduce
                        (fn [updated-frontier coord]
                          (if (not (visited coord))
                            (conj updated-frontier {:coord coord :score (dist coord (:coord target))})
                            updated-frontier))
                        (rest frontier)
                        (in-range-coords indexed current))]
        (recur (conj visited (:coord current))
               (conj trail current)
               (sort-by first frontiers)
               (inc round)))
      trail)))

(defn accessible? [base target indexed]
  (some #(= 1 (:score %)) (get-trails base target indexed)))

(defn reachable-coords [indexed characters character]
  (->> (mapcat (partial in-range-coords indexed)
               (enemies characters character))
       (filter #(accessible? character {:coord %} indexed))))

(defn closest [base coords]
  (->> (map #(do [(dist base %) %]) coords)
       (group-by first)
       (apply min-key first)
       second
       (map second)))

(comment
  (println *e)

  (let [indexed   (merge indexed-board (index-pieces @characters))
        character (first @characters)
        in-range  (in-range-enemies indexed @characters character)
        in-range-coord (mapcat (partial in-range-coords indexed)
                               (enemies @characters character))
        reachable-coord (reachable-coords indexed @characters character)
        closest-coord (closest (:coord character) reachable-coord)
        target (first (reverse (sort closest-coord)))
        reachable-pieces (into-piece [target] \+)
        pieces (concat @characters reachable-pieces)]
    (print-input (with-pieces clean-input pieces)))

  (print-input (with-pieces clean-input (list (last @characters))))
  (print-input (with-pieces clean-input (take 1 @characters)))

  (let [chars [(first @characters) (last @characters)]]
    (print-input
      (with-pieces clean-input
        chars
        (map
          #(assoc % :piece (:score %))
          (accessible?
            (first @characters)
            (last @characters)
            (merge indexed-board (index-pieces [(first @characters) (last @characters)])))))))


  (map identity @characters)
  (println (goblins (sort-by :coord @characters)))
  )

