package icu.aimerny.elix.registry;


import icu.aimerny.elix.consts.IdConst;
import net.fabricmc.fabric.api.item.v1.FabricItemSettings;
import net.minecraft.item.Item;
import net.minecraft.registry.Registries;
import net.minecraft.registry.Registry;

public class ModItem {

    public static final Item OBSIDIAN_INGOT = new Item(new FabricItemSettings());

    public static void register(String id, Item item) {
        Registry.register(Registries.ITEM, id, item);
    }

    public static void init() {
        register(IdConst.OBSIDIAN_INGOT, OBSIDIAN_INGOT);
    }

}
